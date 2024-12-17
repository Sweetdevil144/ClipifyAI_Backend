# Dev.md

## Behavior for Adding Shorts to a User's Profile

When the user generates shorts and selects some, you can add them like this:

```go
user := models.User{ID: 1}
selectedShorts := []models.Short{
    {Src: "video_link_1", StartTime: 10.5, EndTime: 15.0, QualityRanking: ptr(5)},
    {Src: "video_link_2", StartTime: 30.0, EndTime: 45.0, QualityRanking: ptr(4)},
}
user.Shorts = selectedShorts
db.Save(&user)
```
