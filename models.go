package main

type Link struct {
	Id    int64  `json:"id,omitempty";db:"id"`
	Url   string `json:"url";db:"url"`
	Short string `json:"short";db:"short"`
}

type Stat struct {
	Id        int64  `json:"id,omitempty";db:"id"`
	Referer   string `json:"referer";db:"referer"`
	UserAgent string `json:"user_agent";db:"user_agent"`
	Ip        string `json:"ip";db:"ip"`
	LinkId    int64  `json:"link_id";db:"link_id"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
