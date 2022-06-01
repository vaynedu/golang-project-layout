.PHONY: gen wire test build run clean

wire:
	# 依赖注入
	cd model/container/ && wire && cd ..
db:
	# 使用db工具生成go代码
build:
	./build.sh
run:
	# 运行该程序
clean:
	# 清理中间生成文件