package storage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"shipa-gen/src/models"
	"time"
)

func (s *service) CountEventsMetrics(configId string, metrics []string) ([]models.Metric, error) {
	now := time.Now()
	nowYear, nowWeek := now.ISOWeek()
	nowMonth := now.Month()
	nowDay := now.YearDay()
	configTypes := bson.A{}
	for _, metric := range metrics {
		configTypes = append(configTypes, metric)
	}
	pipeline :=
		bson.A{
			bson.D{
				{"$match",
					bson.D{
						{"$and",
							bson.A{
								bson.D{{"configId", configId}},
								bson.D{
									{"$expr",
										bson.D{
											{"$in",
												bson.A{
													"$type",
													configTypes,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			bson.D{
				{"$addFields",
					bson.D{
						{"itemYear", bson.D{{"$year", bson.D{{"date", "$createdAt"}}}}},
						{"itemMonth", bson.D{{"$month", bson.D{{"date", "$createdAt"}}}}},
						{"itemWeek", bson.D{{"$week", bson.D{{"date", "$createdAt"}}}}},
						{"itemDay", bson.D{{"$dayOfYear", bson.D{{"date", "$createdAt"}}}}},
					},
				},
			},
			bson.D{
				{"$addFields",
					bson.D{
						{"thisCountYear",
							bson.D{
								{"$cond",
									bson.D{
										{"if",
											bson.D{
												{"$eq",
													bson.A{
														"$itemYear",
														nowYear,
													},
												},
											},
										},
										{"then", 1},
										{"else", 0},
									},
								},
							},
						},
						{"thisCountDay",
							bson.D{
								{"$cond",
									bson.D{
										{"if",
											bson.D{
												{"$and",
													bson.A{
														bson.D{
															{"$eq",
																bson.A{
																	"$itemYear",
																	nowYear,
																},
															},
														},
														bson.D{
															{"$eq",
																bson.A{
																	"$itemDay",
																	nowDay,
																},
															},
														},
													},
												},
											},
										},
										{"then", 1},
										{"else", 0},
									},
								},
							},
						},
						{"thisCountWeek",
							bson.D{
								{"$cond",
									bson.D{
										{"if",
											bson.D{
												{"$and",
													bson.A{
														bson.D{
															{"$eq",
																bson.A{
																	"$itemYear",
																	nowYear,
																},
															},
														},
														bson.D{
															{"$eq",
																bson.A{
																	"$itemWeek",
																	nowWeek,
																},
															},
														},
													},
												},
											},
										},
										{"then", 1},
										{"else", 0},
									},
								},
							},
						},
						{"thisCountMonth",
							bson.D{
								{"$cond",
									bson.D{
										{"if",
											bson.D{
												{"$and",
													bson.A{
														bson.D{
															{"$eq",
																bson.A{
																	"$itemYear",
																	nowYear,
																},
															},
														},
														bson.D{
															{"$eq",
																bson.A{
																	"$itemMonth",
																	nowMonth,
																},
															},
														},
													},
												},
											},
										},
										{"then", 1},
										{"else", 0},
									},
								},
							},
						},
					},
				},
			},
			bson.D{
				{"$group",
					bson.D{
						{"_id", "$type"},
						{"name", bson.D{{"$first", "$type"}}},
						{"total", bson.D{{"$sum", 1}}},
						{"perDay", bson.D{{"$sum", "$thisCountDay"}}},
						{"perWeek", bson.D{{"$sum", "$thisCountWeek"}}},
						{"perMonth", bson.D{{"$sum", "$thisCountMonth"}}},
						{"perYear", bson.D{{"$sum", "$thisCountYear"}}},
					},
				},
			},
		}

	ctx := context.TODO()
	cursor, err := s.events.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}

	var calculated []models.Metric
	if err = cursor.All(ctx, &calculated); err != nil {
		return nil, err
	}

	return calculated, err
}
