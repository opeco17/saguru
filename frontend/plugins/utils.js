const snakeToCamel = function (p) {
  return p.replace(/_./g, s => s.charAt(1).toUpperCase())
}

const camelToSnake = function (p) {
  return p.replace(/([A-Z])/g, s => '_' + s.charAt(0).toLowerCase())
}

export default (context, inject) => {
    inject('snakeToCamel', snakeToCamel)
    inject('camelToSnake', camelToSnake)
}