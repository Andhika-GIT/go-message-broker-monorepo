"use client";

import {
  Dropzone,
  DropzoneContent,
  DropzoneEmptyState,
} from "@/components/ui/shadcn-io/dropzone";
import { useState } from "react";
import { toast } from "sonner";

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
      accept={{
        'excel/xls': ['.xls'],
        'excel/xlsx': ['.xlsx'],
      }}
      maxSize={1024 * 1024 * 10}
      onDrop={handleDrop}
      onError={(e) => {
        toast.error(e.message);
      }}
      src={files}
    >
      <DropzoneEmptyState />
      <DropzoneContent />
    </Dropzone>
  );
};
