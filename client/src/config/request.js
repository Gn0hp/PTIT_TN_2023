
export const RequestParams = {
  host: 'http://localhost:8400',
  path: {
    register: '/api/v1/auth/register',
    login: '/api/v1/auth/login',
    adminLogin: '/api/v1/admin/login',
    adminGetPendingUsers: '/api/v1/admin/pending-users',
    createElection: '/api/v1/election/create',
    registerCandidate: '/api/v1/voter/register-candidate',

    verify_voter: '/api/v1/admin/verify-voter',
    verify_candidate: '/api/v1/admin/verify-candidate',
    view_candidate: '/api/v1/voter/view-candidate',
    vote: '/api/v1/voter/vote',

    get_post_by_candidate_id: '/api/v1/candidate/get-posts-by-candidate-id',
    create_post: '/api/v1/candidate/post-create',
    delete_post: '/api/v1/candidate/post-delete',
    patch_post: '/api/v1/candidate/post-update',

    election_roles_by_election_id: '/api/v1/election-role/find-by-election-id',
    check_election: '/api/v1/election/check'
  }
}
