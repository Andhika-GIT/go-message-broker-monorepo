"use client"

import { handleFetchResponse, SERVER_BASE_URL_FOR_CLIENT } from "@/lib/helper"
import { Error } from "@/lib/types";
import { Error as ResponseError } from "@/lib/types";

export const UploadUserExcel = async (file: File): Promise<string | undefined> => {
  const BASE_URL = `${SERVER_BASE_URL_FOR_CLIENT}/user/upload`;

  try {

    const formData = new FormData()
    formData.append("file", file)

    const response = await fetch(BASE_URL, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body:formData,
    });

    return await handleFetchResponse(response);
  } catch (e) {
    throw e as Error;
  }
};
