const API_BASE_URL = 'http://ec2-18-180-202-55.ap-northeast-1.compute.amazonaws.com'

export default (context, inject) => {
    inject('API_BASE_URL', API_BASE_URL)
}