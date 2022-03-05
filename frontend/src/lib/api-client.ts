import axios from 'axios';

const client = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
  timeout: 10000,
});

console.log(process.env.NEXT_PUBLIC_API_BASE_URL)

export default client;
