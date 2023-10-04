import Vue from 'vue'
import Router from 'vue-router'
import Authentication from '../components/Authentication.vue'
import VoterHome from '../components/voter/VoterHome.vue'
import CandidateHome from '../components/candidate/CandidateHome.vue'
import AdminHome from '../components/admin/AdminHome.vue'

Vue.use(Router)

export default new Router({
  routes: [
    // {
    //   path: '/wallet',
    //   name: 'ConnectWallet',
    //   component: ConnectWallet
    // },
    // {
    //   path: '/',
    //   name: 'HelloWorld',
    //   component: HelloWorld
    // },
    {
      path: '/',
      name: 'Auth',
      component: Authentication
    },
    {
      path: '/voters/home',
      name: 'VoterHome',
      component: VoterHome
    },
    {
      path: '/candidates/home',
      name: 'CandidateHome',
      component: CandidateHome
    },
    {
      path: '/admin/home',
      name: 'AdminHome',
      component: AdminHome
    }
  ]
})
