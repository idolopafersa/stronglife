import axios from 'axios';

const API_URL = 'https://stronglifeapi.fernandezpablo.es/api';

export  const UserImage = async (id:string) => {
  try {
    console.log('Initiating image request...');
    const response = await axios.get(`${API_URL}/user/getimage`, {
        params: {user_id: id}
    });
    
    return response.data;
  } catch (error) {
    
    throw error;
  }
};

 export default UserImage;