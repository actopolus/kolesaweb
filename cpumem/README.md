# Профилирование CPU и Heap

## Собираем профили
1. go test -bench=. -benchmem -benchtime 2s | tee /var/tmp/b1.bench (Бенчмарк и сохранение)
3. go test -bench=. -benchtime 2s -memprofile /var/tmp/mem1.prof    (Cнять профиль памяти)
4. go tool pprof -alloc_objects /var/tmp/mem1.prof                  (Визуализация профиля памяти)
5. go test -bench=. -benchtime 2s -cpuprofile /var/tmp/cpu1.prof    (Снять профиль по процессору)
6. go tool pprof /var/tmp/cpu1.prof                                 (Визуализация профиля процессора)
7. go tool pprof --nodefraction=0.1 /var/tmp/cpu1.prof              (Убираем шум)


## findUser2
1. Оптимизируем
2. go test -bench=. -benchmem -benchtime 2s | tee /var/tmp/b2.bench (Бенчмарк и сохранение)
3. go test -bench=. -benchtime 2s -cpuprofile /var/tmp/cpu2.prof    (Снять профиль по процессору)
4. go tool pprof --nodefraction=0.1 /var/tmp/cpu2.prof              (Визуализация профиля процессора)
5. benchcmp /var/tmp/b1.bench /var/tmp/b2.bench                     (Сравнение результатов)

## findUser3
1. Оптимизируем
2. go test -bench=. -benchmem -benchtime 2s | tee /var/tmp/b3.bench (Бенчмарк и сохранение)
6. benchcmp /var/tmp/b2.bench /var/tmp/b3.bench                     (Сравнение)