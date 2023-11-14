<template>
  <header>

    <nav class="navbar navbar-dark navbar-expand-lg bg-dark text-white" data-bs-theme="dark">
      <div class="container-fluid">

        <a class="navbar-brand" href="/#/voters/home">Election App</a>

        <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
          <span class="navbar-toggler-icon"></span>
        </button>

        <div class="collapse navbar-collapse"
             id="navbarNav">
          <ul class="navbar-nav">
            <li class="nav-item">
              <a class="nav-link active" aria-current="page" href="/#/voters/home">Home</a>
            </li>
            <li class="nav-item">
              <a class="nav-link" href="#">Features</a>
            </li>
            <li class="nav-item">
              <a class="nav-link" href="#">Contact</a>
            </li>
            <li class="nav-item">
              <a class="nav-link" href="#">Help</a>
            </li>
            <li class="nav-item">
              <a class="nav-link" href="#">About</a>
            </li>
          </ul>
        </div>

        <div class="d-flex ms-auto">
          <span class="navbar-text" style="color: #fff">
            Welcome, {{userName}}
          </span>

         <a href="/#/me/profile">
           <button class="btn btn-outline-info me-2" @click="viewProfile">
             View Profile
           </button>
         </a>

          <button class="btn btn-outline-danger" @click="logout">
            Log Out
          </button>

        </div>

      </div>
    </nav>

  </header>
</template>

<script>
import {AxiosInstance} from '../../config/auth'

export default {
  name: 'Header',
  props: {
    userName: String
  },
  created () {
    const userGlobal = JSON.parse(sessionStorage.getItem('userGlobal'))
    this.userName = userGlobal.full_name
  },
  methods: {
    viewProfile () {
      // view profile logic
    },
    logout () {
      // logout logic
      sessionStorage.removeItem('userGlobal')
      sessionStorage.removeItem('accessToken')
      sessionStorage.removeItem('electionId')
      sessionStorage.removeItem('hasElection')
      AxiosInstance.defaults.headers.common['Authorization'] = null
      this.$router.push('/')
    }
  }
}
</script>
