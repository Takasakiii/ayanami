-- CreateTable
CREATE TABLE "uploaded_files" (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "created_at" DATETIME NOT NULL,
    "user_agent" TEXT NOT NULL,
    "ip" TEXT NOT NULL,
    "file_name" TEXT NOT NULL,
    "mime_type" TEXT NOT NULL
);
