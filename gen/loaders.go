package gen

import (
	"context"

	"github.com/graph-gophers/dataloader"
)

func GetLoaders(db *DB) map[string]*dataloader.Loader {
	loaders := map[string]*dataloader.Loader{}

	companiesBatchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result

		ids := make([]string, len(keys))
		for i, key := range keys {
			ids[i] = key.String()
		}

		items := &[]Company{}
		res := db.Query().Find(items, "id IN (?)", ids)
		if res.Error != nil && !res.RecordNotFound() {
			return []*dataloader.Result{
				&dataloader.Result{Error: res.Error},
			}
		}

		itemMap := make(map[string]Company, len(keys))
		for _, item := range *items {
			itemMap[item.ID] = item
		}

		for _, key := range keys {
			id := key.String()
			item, ok := itemMap[id]
			if !ok {
				results = append(results, &dataloader.Result{
					Data:  nil,
					Error: nil,
					// Error: fmt.Errorf("Company with id '%s' not found", id),
				})
			} else {
				results = append(results, &dataloader.Result{
					Data:  &item,
					Error: nil,
				})
			}
		}
		return results
	}

	loaders["Company"] = dataloader.NewBatchedLoader(companiesBatchFn, dataloader.WithClearCacheOnBatch())

	usersBatchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result

		ids := make([]string, len(keys))
		for i, key := range keys {
			ids[i] = key.String()
		}

		items := &[]User{}
		res := db.Query().Find(items, "id IN (?)", ids)
		if res.Error != nil && !res.RecordNotFound() {
			return []*dataloader.Result{
				&dataloader.Result{Error: res.Error},
			}
		}

		itemMap := make(map[string]User, len(keys))
		for _, item := range *items {
			itemMap[item.ID] = item
		}

		for _, key := range keys {
			id := key.String()
			item, ok := itemMap[id]
			if !ok {
				results = append(results, &dataloader.Result{
					Data:  nil,
					Error: nil,
					// Error: fmt.Errorf("User with id '%s' not found", id),
				})
			} else {
				results = append(results, &dataloader.Result{
					Data:  &item,
					Error: nil,
				})
			}
		}
		return results
	}

	loaders["User"] = dataloader.NewBatchedLoader(usersBatchFn, dataloader.WithClearCacheOnBatch())

	tasksBatchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result

		ids := make([]string, len(keys))
		for i, key := range keys {
			ids[i] = key.String()
		}

		items := &[]Task{}
		res := db.Query().Find(items, "id IN (?)", ids)
		if res.Error != nil && !res.RecordNotFound() {
			return []*dataloader.Result{
				&dataloader.Result{Error: res.Error},
			}
		}

		itemMap := make(map[string]Task, len(keys))
		for _, item := range *items {
			itemMap[item.ID] = item
		}

		for _, key := range keys {
			id := key.String()
			item, ok := itemMap[id]
			if !ok {
				results = append(results, &dataloader.Result{
					Data:  nil,
					Error: nil,
					// Error: fmt.Errorf("Task with id '%s' not found", id),
				})
			} else {
				results = append(results, &dataloader.Result{
					Data:  &item,
					Error: nil,
				})
			}
		}
		return results
	}

	loaders["Task"] = dataloader.NewBatchedLoader(tasksBatchFn, dataloader.WithClearCacheOnBatch())

	return loaders
}
