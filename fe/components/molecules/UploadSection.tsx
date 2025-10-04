"use client";

import React, { useState } from "react";
import { Upload } from "./Upload";
import { Button } from "../ui/button";

type UploadSectionParam = {
  uploadFn : (file: File) => Promise<string | undefined>
}

export const UploadSection : React.FC<UploadSectionParam> = ({uploadFn}) => {
    const [uploadedFile, setUploadedFile] = useState<File[] | undefined>();
    const [isLoading, setIsLoading] = useState<boolean>(false)
    const handleFileChange = (file: File[]) => {
      console.log("file di upload");
      console.log(file);
      setUploadedFile(file)
    };

    const sendFile = async () => {
      if (!uploadedFile) return
      setIsLoading(true)

      try {
        await uploadFn(uploadedFile[0])
      }
      catch(e) {
        console.log(e)
      }
      finally{
        setIsLoading(false)
      }
      

    }
  return (
    <div className="px-6 flex-col w-full">
      <Upload onFilesChange={handleFileChange}></Upload>
      {uploadedFile && <Button disabled={isLoading} onClick={sendFile} className="w-full mt-3 cursor-pointer disabled:opacity-30" variant="outline">{isLoading ? 'Please wait...' : 'Save'}</Button>}
    </div>
  );
};
