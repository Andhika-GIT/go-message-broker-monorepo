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

export const getAllOrders = async (page: number, pageSize: number, search: null | string = null): Promise<FetchResult | undefined> => {
    const params = new URLSearchParams({
    page: page.toString(),
    per_page: pageSize.toString(),
  });

  if (search && search.trim() !== '') {
    params.append('search', search.trim());
  }

  const BASE_URL = `${SERVER_BASE_URL_FOR_CLIENT}/order?${params.toString()}`;

  try {
    const response = await fetch(BASE_URL, {
      method: "GET",
    });

    return await handleFetchResponse(response);
  } catch (e) {
    throw e as Error;
  }
};