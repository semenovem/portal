#### throw Типизированный error
  
  
```
// Создание ошибки
func foo() error {
	return throw.NewBadRequest("test_err").
		SetDesc("описание {%s}", "параметр для шаблонизатора").
		AddTrace("foo", map[string]any{"userID": "68c7482b-e018-4397-8988-2c1ad91a9d9a"}) // добавляем трейс
}

....
err := foo()
if err != nil {
	// добавление к ошибке трейса
	return throw.Cast(err).AddTrace("func1.foo", nil)
}
```
