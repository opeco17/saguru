const API_BASE_URL = 'http://localhost:8000'

export default (context, inject) => {
    inject('API_BASE_URL', API_BASE_URL)
}