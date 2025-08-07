"use client";

import {
  Dropzone,
  DropzoneContent,
  DropzoneEmptyState,
} from "@/components/ui/shadcn-io/dropzone";
import { useState } from "react";

type UploadProps = {
  onFilesChange: (file: File[]) => void;
};

export const Upload: React.FC<UploadProps> = ({ onFilesChange }) => {
  const [files, setFiles] = useState<File[] | undefined>();
  const handleDrop = (files: File[]) => {
    setFiles(files);
    onFilesChange(files);
  };
  return (
    <Dropzone
      maxSize={1024 * 1024 * 10}
      minSize={1024}
      onDrop={handleDrop}
      onError={console.error}
      src={files}
    >
      <DropzoneEmptyState />
      <DropzoneContent />
    </Dropzone>
  );
};
