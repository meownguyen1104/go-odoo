#!/usr/bin/env bash
# Script nâng cao để chạy Odoo generator, hỗ trợ danh sách module động.

# Dừng script ngay khi có lỗi
set -euo pipefail

# Kiểm tra sự tồn tại của tệp .env
if [ ! -f .env ]; then
    echo "❌ Lỗi: Không tìm thấy tệp .env. Vui lòng tạo tệp này."
    exit 1
fi

# Tải các biến môi trường từ file .env
set -a
source .env
set +a

# --- Cấu hình Modules ---



# 2. Đọc các module tùy chỉnh từ .env vào một mảng khác
CUSTOM_MODULES_ARRAY=()
if [ -n "${CUSTOM_MODULES-}" ]; then
    # Đọc chuỗi được phân tách bằng dấu cách vào mảng
    read -r -a CUSTOM_MODULES_ARRAY <<< "$CUSTOM_MODULES"
fi

# 3. Kết hợp hai mảng lại với nhau
ALL_MODULES_ARRAY=("${CUSTOM_MODULES_ARRAY[@]}")

# 4. Chuyển đổi mảng cuối cùng thành một chuỗi duy nhất, phân tách bằng dấu phẩy
# Đây là định dạng mà cờ -m của generator yêu cầu.
ALL_MODULES_STRING=$(IFS=,; echo "${ALL_MODULES_ARRAY[*]}")


# --- Chạy câu lệnh ---
echo "🚀 Bắt đầu chạy generator..."
echo "👤 User: $ODOO_USER"
echo "🌐 Host: $ODOO_HOST"
echo "📦 Modules sẽ được tạo: $ALL_MODULES_STRING"
echo "----------------------------------------"

# Chạy lệnh go với chuỗi module đã được xử lý
go run generator/main.go \
  -u "${ODOO_USER}" \
  -p "${ODOO_PASSWORD}" \
  -d "${ODOO_DB}" \
  -o ./ \
  --url "${ODOO_HOST}" \
  -t ./generator/cmd/tmpl/model.tmpl \
  -m "${ALL_MODULES_STRING}"

echo "✅ Hoàn thành!"