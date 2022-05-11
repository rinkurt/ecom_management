import axios from 'axios'

const request = axios.create({
  baseURL: 'http://localhost:8090',
  timeout: 1000 * 5,
  headers: { 'Authorization': 'Bearer ' + localStorage.token }
})

export default request
