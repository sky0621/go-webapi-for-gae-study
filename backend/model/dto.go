package model

// Dto ... 用途はマーカーインタフェースだがダックタイプのためダミーメソッドを定義
type Dto interface {
	IsDto() bool
}
