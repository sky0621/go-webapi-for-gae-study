<template>
  <v-layout
    column
    justify-center
    wrap
    align-space-between
    class="ma-3"
  >

    <v-text-field
      v-model="email"
      label="Email"
      required
    />

    <v-text-field
      :type="'password'"
      v-model="password"
      label="Password"
      required
    />

    <v-btn
      class="green white--text"
      @click="submit"
    >
      LOGIN
    </v-btn>

  </v-layout>

</template>

<script>
export default {
  layout: 'login',
  asyncData (context) {
    return {
      email: '',
      password: ''
    }
  },
  methods: {
    async submit() {
      await this.$axios
        .$post('http://localhost:8080/api/v1/login', {
          'email': this.email,
          'password': this.password
        })
        .then(res => (alert(res.id)))
    }
  },
}
</script>
