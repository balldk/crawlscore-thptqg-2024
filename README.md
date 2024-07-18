# Crawl toàn bộ điểm thi THPTQG-2024

Script nho nhỏ viết bằng Go để crawl toàn bộ điểm thi THPTQG-2024, được clone và chỉnh sửa lại từ [script của năm 2021](https://github.com/balldk/crawlscore-thptqg-2021).

## Yêu cầu

1. Đã tải source về
2. Đã cài đặt Go

## Cách dùng

Tải Dependencies

```bash
go mod download
```

Chạy chương trình

```bash
go run .
```

hoặc bạn cũng có thể build ra binary

```bash
go build .
./crawlscore
```

## Tuỳ chỉnh tham số

Một số tham số bạn có thể thay đổi trong file `.env`

```env
MAX_THREAD=3000
OUTPUT_FOLDER=data
TOTAL_FILENAME=total.csv
```

-   `MAX_THREAD` là số luồng tối đa, không nên để quá cao vì web sẽ sập đó, tầm 1000 đến 3000 là ổn.
-   `OUTPUT_FOLDER` là tên thư mục mà dữ liệu được xuất ra.
-   `TOTAL_FILENAME` là tên của file tổng hợp tất cả dữ liệu từ 64 tỉnh thành.
