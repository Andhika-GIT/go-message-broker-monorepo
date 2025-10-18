import { handleFetchResponse, SERVER_BASE_URL } from "@/lib/helper";
import { Error } from "@/lib/types";
import { User } from "@/lib/schemas";

export const getAllUsers = async (): Promise<User[] | undefined> => {
  const BASE_URL = `${SERVER_BASE_URL}/user`;

  try {
    const response = await fetch(BASE_URL, {
      method: "GET",
    });

    return await handleFetchResponse(response);
  } catch (e) {
    throw e as Error;
  }
};
