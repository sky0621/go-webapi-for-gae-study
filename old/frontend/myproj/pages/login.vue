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

      console.log(this.email)
      console.log(this.password)

      firebase.auth().signInWithEmailAndPassword(this.email, this.password).catch(function(error) {
        // Handle Errors here.
        var errorCode = error.code;
        var errorMessage = error.message;
        console.log(errorCode)
        console.log(errorMessage)
      })

          const user = firebase.auth().currentUser
          if (user != null) {
            name = user.displayName;
            email = user.email;
            photoUrl = user.photoURL;
            emailVerified = user.emailVerified;
            uid = user.uid;
            userJSON = user.toJSON
            console.log(name)
            console.log(email)
            console.log(photoUrl)
            console.log(emailVerified)
            console.log(uid)
            console.log(userJSON)
            jwt = user.jwt
            console.log(jwt)

            localStorage.setItem('jwt', jwt)

            this.$axios
              .$post('http://localhost:8080/api/v1/login', {'jwt': jwt})
              .then(res => {
                alert(res.id)
                this.$router.push('/')
              })
          } else {
            alert('No res')
          }
    }
  }
}
</script>
