import { PUBLIC_CLIENT_SERVICE_BASE_URL } from "$env/static/public";
import { fail } from "@sveltejs/kit";
import type { Actions } from "./$types";
import type { ClientServiceResponse } from "$lib/types/common";

export const actions = {
  signOut: async (event) => {
    console.log('signout')
    const response = await fetch(`${PUBLIC_CLIENT_SERVICE_BASE_URL}/auth/signout`, {
      method: 'post',
      headers: {
        'Authorization': `Bearer ${event.cookies.get('mikuuuu')}`
      },
      credentials: 'include'
    })
    const responseData: ClientServiceResponse<{}> = await response.json()

    if (!response.ok) {
      return fail(500, { message: "Something went wrong, please try again later" })
    }

    event.cookies.delete('mikuuuu')

    return {
      message: responseData.message
    }
  }
} satisfies Actions