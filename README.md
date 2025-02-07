
# **🛠️ Persiapan Awal**
Kita harus **menginstal Android SDK, Gradle, dan beberapa alat CLI lainnya** secara manual.  

### **🔹 1. Install Java JDK (Wajib untuk Gradle)**
1. Unduh dan install **JDK (Java Development Kit)**:  
   👉 **Download:** [https://openjdk.org/](https://openjdk.org/)
2. Setelah instalasi, cek versi dengan perintah:  
   ```sh
   java -version
   ```
   Jika berhasil, akan muncul output seperti:
   ```
   openjdk version "17.0.13" 2024-10-15
   OpenJDK Runtime Environment OpenLogic-OpenJDK (build 17.0.13+11-adhoc..jdk17u)
   OpenJDK 64-Bit Server VM OpenLogic-OpenJDK (build 17.0.13+11-adhoc..jdk17u, mixed mode, sharing)
   ```

---

### **🔹 2. Install Android SDK (Tanpa Android Studio)**
1. **Unduh Android Command Line Tools:**  
   👉 [https://developer.android.com/studio#command-tools](https://developer.android.com/studio#command-tools)  
   ![image](https://github.com/user-attachments/assets/bfaf3abe-0a10-4bb1-9832-2bf3d4bcc316)  
2. **Ekstrak file ke direktori berikut:**  
   ```sh
   C:\Android\cmdline-tools
   ```
3. **Buat folder untuk SDK:**  
   ```sh
   C:\Android\Sdk
   ```
4. **Tambahkan ke Environment Variables:**  
   - **Buka** Start Menu → cari **"Edit the system environment variables"**  
   - Klik **Environment Variables**  
   - Tambahkan variabel berikut di **System Variables**:
     - **ANDROID_HOME** → `C:\Android\Sdk`
     - **ANDROID_SDK_ROOT** → `C:\Android\Sdk`
     - **Path** → Tambahkan:
       ```
       C:\Android\cmdline-tools\latest\bin
       C:\Android\Sdk\platform-tools
       C:\Android\Sdk\build-tools\33.0.2
       ```
       ![image](https://github.com/user-attachments/assets/72fce67f-ccc0-4aa3-8a77-478b5c937707)  
5. **Buka Command Prompt (cmd) dan jalankan perintah berikut:**  
   ```sh
   sdkmanager --install "platforms;android-33" "build-tools;33.0.2" "platform-tools"
   sdkmanager --licenses
   ```
   > Pilih **Y** untuk menyetujui lisensi.

---

### **🔹 3. Install Gradle (Tanpa Android Studio)**
1. **Unduh Gradle:**  
   👉 [https://gradle.org/releases/](https://gradle.org/releases/)  
2. **Ekstrak ke folder, misalnya:**  
   ```
   C:\Gradle
   ```
3. **Tambahkan ke Environment Variables:**
   - **Path** → Tambahkan:
     ```
     C:\Gradle\bin
     ```
4. **Cek apakah Gradle terinstal dengan benar:**  
   ```sh
   gradle -v
   ```
   Jika berhasil, akan muncul output versi Gradle.
   ```sh
   
   ------------------------------------------------------------
   Gradle 8.12.1
   ------------------------------------------------------------
   
   Build time:    2025-01-24 12:55:12 UTC
   Revision:      0b1ee1ff81d1f4a26574ff4a362ac9180852b140
   
   Kotlin:        2.0.21
   Groovy:        3.0.22
   Ant:           Apache Ant(TM) version 1.10.15 compiled on August 25 2024
   Launcher JVM:  17.0.13 (OpenLogic 17.0.13+11-adhoc..jdk17u)
   Daemon JVM:    C:\Program Files\OpenLogic\jdk-17.0.13.11-hotspot (no JDK specified, using current Java home)
   OS:            Windows 10 10.0 amd64

   ```

---

### **🔹 4. Buat Proyek Web
Karena kita **tidak menggunakan npm**, kita harus membuat proyek Capacitor secara manual.

1. **Buat folder proyek:**
   ```sh
   mkdir C:\MyCapacitorApp
   cd C:\MyCapacitorApp
   ```
2. **Buat folder untuk file web (HTML, CSS, JS):**
   ```sh
   mkdir www
   ```
3. **Buat file `index.html` di dalam `www/`**:
   ```html
   <!DOCTYPE html>
   <html lang="en">
   <head>
       <meta charset="UTF-8">
       <meta name="viewport" content="width=device-width, initial-scale=1.0">
       <title>Capacitor Without npm</title>
       <script type="module">
           import { Device } from 'https://cdn.jsdelivr.net/npm/@capacitor/device@latest/dist/esm/index.js';

           async function getDeviceInfo() {
               const info = await Device.getInfo();
               document.getElementById("output").innerText = JSON.stringify(info, null, 2);
           }
       </script>
   </head>
   <body>
       <h1>Capacitor Without npm</h1>
       <button onclick="getDeviceInfo()">Get Device Info</button>
       <pre id="output"></pre>
   </body>
   </html>
   ```

4. **Buat file konfigurasi `capacitor.config.json`:**
   ```json
   {
     "appId": "com.example.myapp",
     "appName": "My Capacitor App",
     "webDir": "www",
     "bundledWebRuntime": false
   }
   ```

---

### **🔹 5. Buat Proyek Android Secara Manual**
Karena kita **tidak menggunakan Android Studio**, kita harus membuat proyek Android sendiri.

1. **Masuk ke folder proyek:**
   ```sh
   cd C:\MyCapacitorApp
   ```
2. **Buat folder `android/`:**
   ```sh
   mkdir android
   cd android
   ```
3. **Buat file `settings.gradle`**:
   ```gradle
   rootProject.name = 'MyCapacitorApp'
   include ':app'
   ```
4. **Buat file `build.gradle`**:
   ```gradle
   buildscript {
       repositories {
           google()
           mavenCentral()
       }
       dependencies {
           classpath 'com.android.tools.build:gradle:7.0.0'
       }
   }

   allprojects {
       repositories {
           google()
           mavenCentral()
       }
   }
   ```
5. **Buat file `app/build.gradle`**:
   ```gradle
   apply plugin: 'com.android.application'

   android {
       compileSdkVersion 33

       defaultConfig {
           applicationId "com.example.myapp"
           minSdkVersion 21
           targetSdkVersion 33
           versionCode 1
           versionName "1.0"
       }
   }
   ```

---

### **🔹 6. Build APK Tanpa Android Studio**
Sekarang kita akan membangun APK menggunakan **Gradle CLI**.

1. **Masuk ke folder Android:**
   ```sh
   cd C:\MyCapacitorApp\android
   ```
2. **Jalankan build:**
   ```sh
   gradle assembleDebug
   ```
3. **Hasil APK akan ada di:**
   ```
   C:\MyCapacitorApp\android\app\build\outputs\apk\debug\app-debug.apk
   ```

---

### **🔹 7. Install dan Jalankan APK di Perangkat**
Jika Anda ingin menjalankan aplikasi di perangkat tanpa Android Studio:

1. **Hubungkan perangkat Android ke PC**  
   - Aktifkan **USB Debugging** di perangkat.
   - Cek apakah perangkat terdeteksi:
     ```sh
     adb devices
     ```
2. **Install APK ke perangkat:**
   ```sh
   adb install C:\MyCapacitorApp\android\app\build\outputs\apk\debug\app-debug.apk
   ```
3. **Jalankan aplikasi di perangkat:**
   ```sh
   adb shell am start -n com.example.myapp/.MainActivity
   ```

---

## **✅ Kesimpulan**
✔ **Tanpa Android Studio**, tetap bisa pakai Capacitor 🚀  
✔ **Install Android SDK + Gradle manual** untuk build APK  
✔ **Gunakan ADB untuk install dan jalankan APK**  
✔ **Lebih ringan & fleksibel, tanpa dependency npm atau IDE berat**  

👉 **Sekarang Anda bisa membuat aplikasi mobile dengan Capacitor tanpa npm, tanpa npx, dan tanpa Android Studio!** 🎉🚀
