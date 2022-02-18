package converter

import (
	"reflect"
	"testing"
	"time"

	"emwell/internal/core/diary/entites"
)

func TestConverter_ConvertToPoints(t *testing.T) {
	type args struct {
		rawData []entites.EmotionalInfo
	}
	userID := int64(1)
	beforeDate := time.Date(2002, 21, 11, 12, 15, 53, 0, time.UTC)
	date := time.Date(2002, 21, 12, 12, 15, 53, 0, time.UTC)
	afterDate := time.Date(2002, 21, 13, 12, 15, 53, 0, time.UTC)

	tests := []struct {
		name string
		args args
		want []entites.EmotionalInfo
	}{
		{
			name: "empty input",
			args: args{
				rawData: []entites.EmotionalInfo{},
			},
			want: []entites.EmotionalInfo{},
		},
		{
			name: "nil input",
			args: args{
				rawData: []entites.EmotionalInfo{},
			},
			want: []entites.EmotionalInfo{},
		},
		{
			name: "rate is 0",
			args: args{
				rawData: []entites.EmotionalInfo{
					entites.NewEmotionalDiaryEntity(1, date, entites.EmotionalInfo{
						UserID:        userID,
						EmotionalRate: 0,
						ReferToDate:   beforeDate,
					}),
					entites.NewEmotionalDiaryEntity(1, date, entites.EmotionalInfo{
						UserID:        userID,
						EmotionalRate: 0,
						ReferToDate:   date,
					}),
					entites.NewEmotionalDiaryEntity(1, date, entites.EmotionalInfo{
						UserID:        userID,
						EmotionalRate: 0,
						ReferToDate:   afterDate,
					}),
				},
			},
			want: []entites.EmotionalInfo{},
		},
		{
			name: "1 rate for each day",
			args: args{
				rawData: []entites.EmotionalInfo{
					entites.NewEmotionalDiaryEntity(1, date, entites.EmotionalInfo{
						UserID:        userID,
						EmotionalRate: 100,
						ReferToDate:   beforeDate,
					}),
					entites.NewEmotionalDiaryEntity(1, date, entites.EmotionalInfo{
						UserID:        userID,
						EmotionalRate: 200,
						ReferToDate:   date,
					}),
					entites.NewEmotionalDiaryEntity(1, date, entites.EmotionalInfo{
						UserID:        userID,
						EmotionalRate: 300,
						ReferToDate:   afterDate,
					}),
				},
			},
			want: []entites.EmotionalInfo{
				entites.NewEmotionalDiaryEntity(1, date, entites.EmotionalInfo{
					UserID:        userID,
					EmotionalRate: 100,
					ReferToDate:   beforeDate,
				}),
				entites.NewEmotionalDiaryEntity(1, date, entites.EmotionalInfo{
					UserID:        userID,
					EmotionalRate: 200,
					ReferToDate:   date,
				}),
				entites.NewEmotionalDiaryEntity(1, date, entites.EmotionalInfo{
					UserID:        userID,
					EmotionalRate: 300,
					ReferToDate:   afterDate,
				}),
			},
		},
		{
			name: "1 rate for each day except one in middle",
			args: args{
				rawData: []entites.EmotionalInfo{
					entites.NewEmotionalDiaryEntity(1, date, entites.EmotionalInfo{
						UserID:        userID,
						EmotionalRate: 100,
						ReferToDate:   beforeDate,
					}),
					entites.NewEmotionalDiaryEntity(1, date, entites.EmotionalInfo{
						UserID:        userID,
						EmotionalRate: 0,
						ReferToDate:   date,
					}),
					entites.NewEmotionalDiaryEntity(1, date, entites.EmotionalInfo{
						UserID:        userID,
						EmotionalRate: 300,
						ReferToDate:   afterDate,
					}),
				},
			},
			want: []entites.EmotionalInfo{
				entites.NewEmotionalDiaryEntity(1, date, entites.EmotionalInfo{
					UserID:        userID,
					EmotionalRate: 100,
					ReferToDate:   beforeDate,
				}),
				entites.NewEmotionalDiaryEntity(1, date, entites.EmotionalInfo{
					UserID:        userID,
					EmotionalRate: 300,
					ReferToDate:   afterDate,
				}),
			},
		},
		{
			name: "1 rate for each day except one in the begging",
			args: args{
				rawData: []entites.EmotionalInfo{
					entites.NewEmotionalDiaryEntity(1, date, entites.EmotionalInfo{
						UserID:        userID,
						EmotionalRate: 0,
						ReferToDate:   beforeDate,
					}),
					entites.NewEmotionalDiaryEntity(1, date, entites.EmotionalInfo{
						UserID:        userID,
						EmotionalRate: 200,
						ReferToDate:   date,
					}),
					entites.NewEmotionalDiaryEntity(1, date, entites.EmotionalInfo{
						UserID:        userID,
						EmotionalRate: 300,
						ReferToDate:   afterDate,
					}),
				},
			},
			want: []entites.EmotionalInfo{
				entites.NewEmotionalDiaryEntity(1, date, entites.EmotionalInfo{
					UserID:        userID,
					EmotionalRate: 200,
					ReferToDate:   date,
				}),
				entites.NewEmotionalDiaryEntity(1, date, entites.EmotionalInfo{
					UserID:        userID,
					EmotionalRate: 300,
					ReferToDate:   afterDate,
				}),
			},
		},
		{
			name: "1 rate for each day except one in the end",
			args: args{
				rawData: []entites.EmotionalInfo{
					entites.NewEmotionalDiaryEntity(1, date, entites.EmotionalInfo{
						UserID:        userID,
						EmotionalRate: 100,
						ReferToDate:   beforeDate,
					}),
					entites.NewEmotionalDiaryEntity(1, date, entites.EmotionalInfo{
						UserID:        userID,
						EmotionalRate: 200,
						ReferToDate:   date,
					}),
					entites.NewEmotionalDiaryEntity(1, date, entites.EmotionalInfo{
						UserID:        userID,
						EmotionalRate: 0,
						ReferToDate:   afterDate,
					}),
				},
			},
			want: []entites.EmotionalInfo{
				entites.NewEmotionalDiaryEntity(1, date, entites.EmotionalInfo{
					UserID:        userID,
					EmotionalRate: 100,
					ReferToDate:   beforeDate,
				}),
				entites.NewEmotionalDiaryEntity(1, date, entites.EmotionalInfo{
					UserID:        userID,
					EmotionalRate: 200,
					ReferToDate:   date,
				}),
			},
		},
		{
			name: "3 rate in one day",
			args: args{
				rawData: []entites.EmotionalInfo{
					entites.NewEmotionalDiaryEntity(1, date, entites.EmotionalInfo{
						UserID:        userID,
						EmotionalRate: 100,
						ReferToDate:   date,
					}),
					entites.NewEmotionalDiaryEntity(1, date, entites.EmotionalInfo{
						UserID:        userID,
						EmotionalRate: 100,
						ReferToDate:   date,
					}),
					entites.NewEmotionalDiaryEntity(1, date, entites.EmotionalInfo{
						UserID:        userID,
						EmotionalRate: 100,
						ReferToDate:   date,
					}),
				},
			},
			want: []entites.EmotionalInfo{
				entites.NewEmotionalDiaryEntity(1, date, entites.EmotionalInfo{
					UserID:        userID,
					EmotionalRate: 100,
					ReferToDate:   date,
				}),
			},
		},
		{
			name: "2 rate in one day in the beginning",
			args: args{
				rawData: []entites.EmotionalInfo{
					entites.NewEmotionalDiaryEntity(1, date, entites.EmotionalInfo{
						UserID:        userID,
						EmotionalRate: 100,
						ReferToDate:   beforeDate,
					}),
					entites.NewEmotionalDiaryEntity(1, date, entites.EmotionalInfo{
						UserID:        userID,
						EmotionalRate: 100,
						ReferToDate:   beforeDate,
					}),
					entites.NewEmotionalDiaryEntity(1, date, entites.EmotionalInfo{
						UserID:        userID,
						EmotionalRate: 200,
						ReferToDate:   date,
					}),
					entites.NewEmotionalDiaryEntity(1, date, entites.EmotionalInfo{
						UserID:        userID,
						EmotionalRate: 300,
						ReferToDate:   afterDate,
					}),
				},
			},
			want: []entites.EmotionalInfo{
				entites.NewEmotionalDiaryEntity(1, date, entites.EmotionalInfo{
					UserID:        userID,
					EmotionalRate: 100,
					ReferToDate:   beforeDate,
				}),
				entites.NewEmotionalDiaryEntity(1, date, entites.EmotionalInfo{
					UserID:        userID,
					EmotionalRate: 200,
					ReferToDate:   date,
				}),
				entites.NewEmotionalDiaryEntity(1, date, entites.EmotionalInfo{
					UserID:        userID,
					EmotionalRate: 300,
					ReferToDate:   afterDate,
				}),
			},
		},
		{
			name: "2 rate in one day in the middle",
			args: args{
				rawData: []entites.EmotionalInfo{
					entites.NewEmotionalDiaryEntity(1, date, entites.EmotionalInfo{
						UserID:        userID,
						EmotionalRate: 100,
						ReferToDate:   beforeDate,
					}),
					entites.NewEmotionalDiaryEntity(1, date, entites.EmotionalInfo{
						UserID:        userID,
						EmotionalRate: 200,
						ReferToDate:   date,
					}),
					entites.NewEmotionalDiaryEntity(1, date, entites.EmotionalInfo{
						UserID:        userID,
						EmotionalRate: 200,
						ReferToDate:   date,
					}),
					entites.NewEmotionalDiaryEntity(1, date, entites.EmotionalInfo{
						UserID:        userID,
						EmotionalRate: 300,
						ReferToDate:   afterDate,
					}),
				},
			},
			want: []entites.EmotionalInfo{
				entites.NewEmotionalDiaryEntity(1, date, entites.EmotionalInfo{
					UserID:        userID,
					EmotionalRate: 100,
					ReferToDate:   beforeDate,
				}),
				entites.NewEmotionalDiaryEntity(1, date, entites.EmotionalInfo{
					UserID:        userID,
					EmotionalRate: 200,
					ReferToDate:   date,
				}),
				entites.NewEmotionalDiaryEntity(1, date, entites.EmotionalInfo{
					UserID:        userID,
					EmotionalRate: 300,
					ReferToDate:   afterDate,
				}),
			},
		},
		{
			name: "2 rate in one day in the end",
			args: args{
				rawData: []entites.EmotionalInfo{
					entites.NewEmotionalDiaryEntity(1, date, entites.EmotionalInfo{
						UserID:        userID,
						EmotionalRate: 100,
						ReferToDate:   beforeDate,
					}),
					entites.NewEmotionalDiaryEntity(1, date, entites.EmotionalInfo{
						UserID:        userID,
						EmotionalRate: 200,
						ReferToDate:   date,
					}),
					entites.NewEmotionalDiaryEntity(1, date, entites.EmotionalInfo{
						UserID:        userID,
						EmotionalRate: 300,
						ReferToDate:   afterDate,
					}),
					entites.NewEmotionalDiaryEntity(1, date, entites.EmotionalInfo{
						UserID:        userID,
						EmotionalRate: 300,
						ReferToDate:   afterDate,
					}),
				},
			},
			want: []entites.EmotionalInfo{
				entites.NewEmotionalDiaryEntity(1, date, entites.EmotionalInfo{
					UserID:        userID,
					EmotionalRate: 100,
					ReferToDate:   beforeDate,
				}),
				entites.NewEmotionalDiaryEntity(1, date, entites.EmotionalInfo{
					UserID:        userID,
					EmotionalRate: 200,
					ReferToDate:   date,
				}),
				entites.NewEmotionalDiaryEntity(1, date, entites.EmotionalInfo{
					UserID:        userID,
					EmotionalRate: 300,
					ReferToDate:   afterDate,
				}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Converter{}
			if got := c.ConvertToPoints(tt.args.rawData); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertToPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
