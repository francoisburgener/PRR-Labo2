set /a i = %1 - 1
FOR /l %%G IN (0,1,%i%) DO pause && start cmd /k go run main.go -proc %%G -N %1

