import axios from 'axios';
const API_URL = 'http://localhost:8000/api';

export const uploadFileTos3 = async (file: File): Promise<any> => {
    const formData = new FormData();
    formData.append("file", file);
  
    const response = await axios.post(`${API_URL}/upload`, formData, {
      headers: {
        "Content-Type": "multipart/form-data",
      },
    });
  
    return response;
  };
  