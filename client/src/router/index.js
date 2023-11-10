import Vue from 'vue'
import Router from 'vue-router'
import Authentication from '../components/Authentication.vue'
import VoterHome from '../components/voter/VoterHome.vue'
import CandidateHome from '../components/candidate/CandidateHome.vue'
import AdminHome from '../components/admin/AdminHome.vue'
import AdminAuth from '../components/admin/AdminAuth.vue'
import ViewCandidate from '../components/voter/ViewCandidate.vue'
import CandidateDetails from '../components/voter/CandidateDetails.vue'
import RegisterCandidate from '../components/candidate/RegisterCandidate'
import VoteNow from '../components/voter/VoteNow'
import Profile from '../components/Profile'
import VerifyUser from '../components/admin/VerifyUser.vue'
import AdminStat from '../components/admin/Stat.vue'
import AdminStatRoleElect from '../components/admin/StatRoleElect.vue'
import AdminStatBallot from '../components/admin/StatBallot.vue'
import UserManagement from '../components/admin/UserManagement.vue'
import CreateElection from '../components/admin/CreateElection.vue'
import ViewResult from '../components/voter/ViewResult.vue'

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
      path: '/voters/view_candidates/:id',
      name: 'CandidateDetails',
      component: CandidateDetails
    },
    {
      path: '/voters/view_candidates',
      name: 'ViewCandidate',
      component: ViewCandidate
    },
    {
      path: '/voters/vote_now',
      name: 'VoteNow',
      component: VoteNow
    },
    {
      path: '/candidates/home',
      name: 'CandidateHome',
      component: CandidateHome
    },
    {
      path: '/register_candidate',
      name: 'RegisterCandidate',
      component: RegisterCandidate
    },
    {
      path: '/admin/signin',
      name: 'AdminAuthentication',
      component: AdminAuth
    },
    {
      path: '/admin/home',
      name: 'AdminHome',
      component: AdminHome
    },
    {
      path: '/admin/verify',
      name: 'VerifyUSer',
      component: VerifyUser
    },
    {
      path: '/admin/create_election',
      name: 'CreateElection',
      component: CreateElection
    },
    {
      path: '/admin/user-management',
      name: 'UserManagement',
      component: UserManagement
    },
    {
      path: '/admin/stat',
      name: 'AdminStat',
      component: AdminStat
    },
    {
      path: '/admin/stat-role-elect/:id',
      name: 'AdminStatRoleElect',
      component: AdminStatRoleElect
    },
    {
      path: '/admin/stat-ballot/:id',
      name: 'AdminStatBallot',
      component: AdminStatBallot
    },
    {
      path: '/me/profile',
      name: 'Profile',
      component: Profile
    },
    {
      path: '/view_result',
      name: 'ViewResult',
      component: ViewResult
    },
    {
      path: '*',
      redirect: '/'
    }
  ]
})
