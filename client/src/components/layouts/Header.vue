<template>
  <header>

    <nav class="navbar navbar-expand-lg bg-dark" data-bs-theme="dark">
      <div class="container-fluid">

        <a class="navbar-brand" href="#">Election App</a>

        <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
          <span class="navbar-toggler-icon"></span>
        </button>

        <div class="collapse navbar-collapse"
             id="navbarNav">
          <ul class="navbar-nav">
            <li class="nav-item">
              <a class="nav-link active" aria-current="page" href="#">Home</a>
            </li>
            <li class="nav-item">
              <a class="nav-link" href="#">Features</a>
            </li>
            <li class="nav-item">
              <a class="nav-link" href="#">Pricing</a>
            </li>
            <li class="nav-item">
              <a class="nav-link disabled" aria-disabled="true">Disabled</a>
            </li>
          </ul>
        </div>

        <div class="d-flex ms-auto">
          <span class="navbar-text" style="color: #fff">
            Welcome, {{userName}}
          </span>

          <button class="btn btn-outline-info me-2" @click="viewProfile">
            View Profile
          </button>

          <button class="btn btn-outline-danger" @click="logout">
            Log Out
          </button>

        </div>

      </div>
    </nav>

  </header>
</template>

<script>
export default {
  name: 'Header',
  props: {
    userName: String
  },
  methods: {
    viewProfile () {
      // view profile logic
    },
    logout () {
      // logout logic
    }
  }
}
</script>
