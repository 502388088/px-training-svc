package model

type Profile struct {
    Id int `json:"id,omitempty"  db:"id"`
    Eid int `json:"eid,omitempty"  db:"eid"`
    Fullname string `json:"fullname,omitempty"  db:"fullname"`
    Position string `json:"position,omitempty"  db:"position"`
}
