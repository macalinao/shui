clean:
	rm ./shui

shui:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o shui .

docker: clean shui
	docker build .

.PHONY: clean docker