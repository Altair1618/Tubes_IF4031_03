import { RSA_PUBLIC_KEY } from "$env/static/private";
import { redirect, type Handle } from "@sveltejs/kit";
import jwt from 'jsonwebtoken'


export const handle: Handle = async ({ event, resolve }) => {
  const accessToken = event.cookies.get('mikuuuu')
  console.log('baka')

  let authenticated = true
  if (!accessToken) authenticated = false
  else {
    try {
      const decodedToken = jwt.verify(accessToken, RSA_PUBLIC_KEY)
      event.locals.user = { userId: (decodedToken as jwt.JwtPayload)["userId"] }
    } catch (e) {
      console.log(e)
      authenticated = false
    }
  }

  if (!authenticated) {
    if (event.url.pathname.startsWith("/profile")) {
      throw redirect(302, `/signin?redirectTo=${event.url.pathname}`)
    }
  }
  
  const response = await resolve(event)
  return response
}