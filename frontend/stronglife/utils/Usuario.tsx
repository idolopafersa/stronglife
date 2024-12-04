const url: string = 'https://stronglifeapi.fernandezpablo.es';
//Como se importa el asinc storage para guardar el nombre de usuario
export const Login = async (usuario: string, password: string) => {
 try{
    const response = await fetch(`${url}/api/login`, {
    method: 'POST',
    headers: {
      
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ usuario, password }),
  });
  const data = await response.json();
  return data;
}catch(e){
    console.error(e);
    throw new Error('No se ha podido iniciar sesión');
}
}
export const Register = async (usuario: string, password: string) => {
    try{
       const response = await fetch(`${url}/api/user/register`, {
       method: 'POST',
       headers: {
         
         'Content-Type': 'application/json',
       },
       body: JSON.stringify({ usuario, password }),
     });
     const data = await response.json();
     return data;
   }catch(e){
       console.error(e);
       throw new Error('No se ha podido registrar');
   }
   }