new Vue({
  el: '#app',
  data () {
    return {
      data: null
    }
  },
  mounted () {
    axios
      .get('go-backend/v1/')
      .then(response => (this.data = response))
  }
})