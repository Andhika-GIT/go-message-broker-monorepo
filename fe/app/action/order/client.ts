"use client"

import { handleFetchResponse, SERVER_BASE_URL_FOR_CLIENT } from "@/lib/helper"
import { Order } from "@/lib/schemas";
import { Error, Paginate } from "@/lib/types";

type FetchResult = Paginate<Order[]>

export const UploadOrderExcel = async (file: File): Promise<string | undefined> => {
  const BASE_URL = `${SERVER_BASE_URL_FOR_CLIENT}/order/upload`;

  try {

    const formData = new FormData()
    formData.append("file", file)

    const response = await fetch(BASE_URL, {
      method: "POST",
      body:formData,
    });

    return await handleFetchResponse(response);
  } catch (e) {
    throw e as Error;
  }
};

export const getAllOrders = async (page: number, pageSize: number): Promise<FetchResult | undefined> => {
  const BASE_URL = `${SERVER_BASE_URL_FOR_CLIENT}/order?page=${page}&per_page=${pageSize}`;

  try {
    const response = await fetch(BASE_URL, {
      method: "GET",
    });

    return await handleFetchResponse(response);
  } catch (e) {
    throw e as Error;
  }
};