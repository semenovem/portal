### portal 

1. .Info(err.Error()) - рефакторинг переделать приемный тип на Error 
1. сервер для audit (internal/app/audit/grpc-server.go)
2. клиент grpc для отправки в audit (providers/audit)
3. переделать получение времени из конфига - получать через методы (как в config.Audit)
4. в аудит добавить БД clickhouse
- использовать vault для кредов
- настроить метрики 
- настроить графану
- метод /v1/auth/refresh
- 2-й фактор (токен + телефон) 
- мидлваре для контроля корректности установки refresh токена - что бы отправлялся только на указанные api
- 
- 
- 
