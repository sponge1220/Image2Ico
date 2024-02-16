### 在 Windows 上使用`image2Ico`

用戶可以透過命令行介面（Command Line Interface, CLI）來使用這個在 Windows 上運行的圖片轉換程式。以下是使用這個程式的步驟說明：

### 步驟 1: 打開命令提示符或PowerShell
首先，Windows用戶需要打開命令提示符（cmd）或PowerShell。這可以通過在Windows搜尋欄中輸入`cmd`或`PowerShell`，然後選擇對應的應用程式來完成。

### 步驟 2: 導航到程式所在的目錄
使用`cd`命令來改變當前目錄，導航到包含`image2Ico.exe`的目錄。例如，如果`image2Ico.exe`位於`C:\Programs\Image2Ico`目錄中，則輸入：

```cmd
cd C:\Programs\Image2Ico
```

### 步驟 3: 執行程式
在命令行中輸入程式名稱和必要的參數來執行程式。您需要指定要轉換圖片的目錄路徑，並且可以選擇指定輸出ICO的大小。例如：

```cmd
image2Ico.exe -dir="C:\path\to\your\images" -size=256
```

這個命令會將指定目錄`C:\path\to\your\images`中的所有PNG或JPG/JPEG圖片轉換成256x256像素的ICO檔案，並將它們存儲在同一目錄下的`output`資料夾中。

### 可用的命令行參數
- `-dir`: 必須指定的參數，用來指出包含圖片的目錄路徑。
- `-size`: 可選參數，用來指定輸出ICO的大小（像素）。如果不指定，預設值為256x256。

### 注意事項
- 確保指定的目錄路徑是正確的，並且該目錄具有可讀寫的權限。
- 如果`output`資料夾不存在，程式會自動創建它。
- 命令行中的路徑可能需要根據您的系統配置和檔案位置進行調整。

通過這些步驟，用戶就可以輕鬆地使用這個圖片轉換工具了。

為了在Mac和Linux系統上使用`image2Ico`程式，用戶需要通過終端（Terminal）執行程式。以下是針對Mac和Linux用戶的詳細使用教學。

### 在 Mac 上使用`image2Ico`

1. **打開終端（Terminal）**:
    - 可以通過Finder中的應用程式 > 實用工具 > 終端，或者使用Spotlight搜索來打開終端。

2. **導航到程式所在的目錄**:
    - 使用`cd`命令更改到包含`image2Ico`程式的目錄。例如，如果程式位於`Downloads`文件夾，輸入：
      ```bash
      cd ~/Downloads
      ```

3. **賦予程式執行權限** (如果還未設置):
    - 在首次運行之前，可能需要為程式文件設置執行權限。使用`chmod`命令來賦予它執行權限：
      ```bash
      chmod +x image2Ico
      ```

4. **執行程式**:
    - 使用以下命令來運行`image2Ico`程式，其中`-dir`參數指定包含圖片的目錄，`-size`參數指定輸出ICO的大小（這是可選的）：
      ```bash
      ./image2Ico -dir="/path/to/your/images" -size=256
      ```
    - 確保替換`/path/to/your/images`為你的圖片目錄的實際路徑。

### 在 Linux 上使用`image2Ico`

使用步驟與Mac類似，以下是針對Linux用戶的指南：

1. **打開終端（Terminal）**:
    - 通常可以通過系統的應用程式菜單或使用快捷鍵（如Ctrl+Alt+T）來打開終端。

2. **導航到程式所在的目錄**:
    - 使用`cd`命令更改到包含`image2Ico`程式的目錄。

3. **賦予程式執行權限** (如果還未設置):
    - 使用`chmod`命令來賦予程式執行權限：
      ```bash
      chmod +x image2Ico
      ```

4. **執行程式**:
    - 使用以下命令來運行`image2Ico`程式：
      ```bash
      ./image2Ico -dir="/path/to/your/images" -size=256
      ```
    - 記得替換`/path/to/your/images`為你的圖片目錄的實際路徑。
