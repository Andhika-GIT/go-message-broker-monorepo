"use client";

import React, { useState } from "react";
import { Upload } from "./Upload";
import { Button } from "../ui/button";

export const UploadSection = () => {
    const [uploadedFile, setUploadedFile] = useState<File[] | undefined>();
    const handleFileChange = (file: File[]) => {
    console.log("file di upload");
    console.log(file);
    setUploadedFile(file)
  };
  return (
    <div className="px-6 flex-col w-full">
      <Upload onFilesChange={handleFileChange}></Upload>
      {uploadedFile && <Button className="w-full mt-3 cursor-pointer" variant="outline">Save</Button>}
    </div>
  );
};
