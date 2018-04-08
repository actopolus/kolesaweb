# Профилирование CPU и Heap

## Собираем профили
1. go test -bench=. -benchmem -benchtime 2s (Показатели benchmark) 
2. go test -bench=. -benchmem -benchtime 2s \> b1.bench (Сохранить бэнчмарк)
3. go test -bench=. -benchtime 2s -memprofile mem.prof (Cнять профиль памяти)
4. go test -bench=. -benchtime 2s -cpuprofile cpu.prof (Снять профиль по процессору)
5. go tool pprof -alloc_objects mem.prof (Визуализация профиля памяти)
6. go tool pprof cpu.prof (Визуализация профиля процессора)
7. go tool pprof --nodefraction=0.1 cpu.prof (Убираем шум)


## findUser2
1. Оптимизируем
2. go test -bench=. -benchmem -benchtime 2s \> b2.bench (Сохранить бэнчмарк)
3. go test -bench=. -benchtime 2s -memprofile mem.prof (Cнять профиль памяти)
4. go test -bench=. -benchtime 2s -cpuprofile cpu.prof (Снять профиль по процессору)
5. go tool pprof --nodefraction=0.1 cpu.prof (Визуализация профиля процессора)
6. benchcmp b1.bench b2.bench

## findUser3
1. Оптимизируем
2. go test -bench=. -benchmem -benchtime 2s \> b3.bench (Сохранить бэнчмарк)
3. go test -bench=. -benchtime 2s -memprofile mem.prof (Cнять профиль памяти)
4. go tool pprof -alloc_objects mem.prof (Визуализация профиля памяти)
6. benchcmp b2.bench b3.bench