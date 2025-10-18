import { handleFetchResponse, SERVER_BASE_URL_FOR_CLIENT } from "@/lib/helper";
import { Error, User } from "@/lib/types";

export const GetAllUsers = async (): Promise<User[] | undefined> => {
  const BASE_URL = `${SERVER_BASE_URL_FOR_CLIENT}/user`;

  try {
    const response = await fetch(BASE_URL, {
      method: "GET",
    });

    return await handleFetchResponse(response);
  } catch (e) {
    throw e as Error;
  }
};
