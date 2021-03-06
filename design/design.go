package design

import (
	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
)

// API 定義
var _ = API("members", func() {
	// API の説明（タイトルと説明）
	Title("Coinspace Members")
	Description("メンバーズサイト")

	cors.Origin("/.*localhost.*/", func() {
		cors.Headers("content-type")
		cors.Methods("GET", "POST", "PUT", "DELETE", "OPTIONS")
		cors.MaxAge(600)
	})

	// サーバ定義
	Server("members", func() {
		Host("localhost", func() {
			URI("http://0.0.0.0:8000") // HTTP REST API
			URI("grpc://0.0.0.0:8080") // gRPC
		})
	})
})

// サービス定義
var _ = Service("members", func() {
	// 説明
	Description("メンバーの一覧を取得する")

	// メソッド (HTTPでいうところのエンドポントに相当)
	Method("add", func() {
		// ペイロード定義
		Payload(func() {
			// 整数型の属性 `a` これは左の被演算子
			Attribute("a", Int, func() {
				Description("Left operand") // 説明
				Meta("rpc:tag", "1")        // gRPC 用のメタ情報。タグ定義
			})
			// 整数型の属性 `b` これは右の被演算子
			Attribute("b", Int, func() {
				Description("Right operand") // 説明
				Meta("rpc:tag", "2")         // gRPC 用のメタ情報。タグ定義
			})
			Required("a", "b") // a と b は required な属性であることの指定
		})

		Result(Int) // メソッドの返値（整数を返す）

		// HTTP トランスポート用の定義
		HTTP(func() {
			GET("/add/{a}/{b}") // GET エンドポイント
			Response(StatusOK)  // レスポンスのステータスは Status OK = 200 を返す
		})

		// GRPC トランスポート用の定義
		GRPC(func() {
			Response(CodeOK) // 手すぽん巣のステータスは CodeOK を返す
		})
	})
})
