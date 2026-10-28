package service

import (
	"fmt"
	"meticuly-assignment-api/internal/constant"
	"meticuly-assignment-api/internal/model"
	"sort"
	"time"
)

func (s *Service) GetOrders() []model.ProductionOrder {
	orders := s.repo.GetOrders()

	sort.Slice(orders, func(i, j int) bool {
		return orders[i].CreatedAt.After(orders[j].CreatedAt)
	})

	return orders

}

func (s *Service) GetOrderSummary() []model.OrderSummary {
	return s.repo.GetOrderSummary()
}
func (s *Service) CreateOrder(
	order model.CreateOrderRequest,
) (model.ProductionOrder, error) {
	if len(order.Operations) == 0 {
		return model.ProductionOrder{}, fmt.Errorf(
			"at least one operation is required",
		)
	}

	// 1. Resolve operation templates and aggregate requirements.
	materialRequirements := make(map[string]int)
	resourceRequirements := make(map[string]bool)

	for _, operation := range order.Operations {
		template, ok := constant.StepCatalog[operation]
		if !ok {
			return model.ProductionOrder{}, fmt.Errorf(
				"unsupported operation: %s", operation,
			)
		}

		for _, material := range template.Materials {
			if material.Denominator <= 0 {
				return model.ProductionOrder{}, fmt.Errorf(
					"invalid material ratio for %s", material.Name,
				)
			}

			// Reserve whole integer quantities.
			materialRequirements[material.Name] += material.Numerator
		}

		for _, resourceName := range template.Resources {
			resourceRequirements[resourceName] = true
		}
	}

	// 2. Validate material availability.
	materials := s.repo.GetMaterialStock()
	materialByName := make(map[string]model.MaterialStock, len(materials))

	for _, material := range materials {
		materialByName[material.Name] = material
	}

	for name, required := range materialRequirements {
		material, ok := materialByName[name]
		if !ok {
			return model.ProductionOrder{}, fmt.Errorf(
				"material not found: %s", name,
			)
		}

		available := material.Stock - material.Consumed - material.Reserved
		if available < float64(required) {
			return model.ProductionOrder{}, fmt.Errorf(
				"insufficient stock for %s: required %d %s, available %.3f %s",
				name,
				required,
				material.Unit,
				available,
				material.Unit,
			)
		}
	}

	// 3. Validate resource availability.
	resources := s.repo.GetResouces()
	resourceByName := make(map[string]model.Resource, len(resources))

	for _, resource := range resources {
		resourceByName[resource.Name] = resource
	}

	for name := range resourceRequirements {
		resource, exists := resourceByName[name]
		if !exists {
			return model.ProductionOrder{}, fmt.Errorf(
				"resource not found: %s", name,
			)
		}

		if resource.Status == model.ResourceStatusBlocked {
			return model.ProductionOrder{}, fmt.Errorf(
				"resource %s is not available (status: %s)",
				name,
				resource.Status,
			)
		}
	}

	// 4. Create the order after validation passes.
	created, err := s.repo.CreateOrder(order)
	if err != nil {
		return model.ProductionOrder{}, fmt.Errorf(
			"create production order: %w", err,
		)
	}

	// 5. Reserve materials.
	for name, required := range materialRequirements {
		material := materialByName[name]
		material.Reserved += float64(required)

		if err := s.repo.UpdateMaterialStock(material); err != nil {
			return model.ProductionOrder{}, fmt.Errorf(
				"reserve material %s: %w", name, err,
			)
		}
	}

	// 6. Associate resources with the order.
	for name := range resourceRequirements {
		resource := resourceByName[name]
		resource.Orders = append(resource.Orders, created.ID)

		if err := s.repo.UpdateResource(resource); err != nil {
			return model.ProductionOrder{}, fmt.Errorf(
				"assign resource %s: %w", name, err,
			)
		}
	}

	// 7. Write creation log.
	s.repo.WriteLog(model.Log{
		OrderID: created.ID,
		Message: fmt.Sprintf(
			"Production order %s created", created.ID,
		),
		Level: "info",
		Time:  time.Now().Format("15:04:05"),
	})

	return created, nil
}