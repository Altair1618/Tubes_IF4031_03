import { PUBLIC_CLIENT_SERVICE_BASE_URL } from "$env/static/public";
import { error } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";
import type { ClientServiceResponse } from "$lib/types/common";

interface UserData {
  email: string;
  picture: string;
  name: string;
  id: string;
}

export const load: PageServerLoad = async ({ params, cookies }) => {
  const response = await fetch(`${PUBLIC_CLIENT_SERVICE_BASE_URL}/profile`, {
    method: 'GET',
    headers: {
      'Authorization': `Bearer ${cookies.get('mikuuuu')}`
    },
    credentials: "include"
  })
  const responseData: ClientServiceResponse<UserData>= await response.json()


  if (!response.ok) {
    throw error(404, 'Not found')
  }

  return responseData
}