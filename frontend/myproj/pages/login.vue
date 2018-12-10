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
import firebase from 'firebase'
export default {
  layout: 'login',
  asyncData (context) {
    return {
      email: '',
      password: ''
    }
  },
  methods: {
    submit() {
      const config = {
        // https://console.firebase.google.com/u/0/project/[project-id]/authentication/users
      }
      firebase.initializeApp(config)

      firebase.auth().signInWithEmailAndPassword(this.email, this.password).then(res => { 
        localStorage.setItem('jwt', res.user.qa)

        this.$axios
          .$post('http://localhost:8080/api/v1/login', {
            'jwt': res.user.qa
          })
          .then(res => (alert(res.id)))

        this.$router.push('/')
      })
    }
  },
}
</script>
