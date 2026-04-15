import client from './client'
import type { ApiResponse, PaginatedData } from '@/types/api'
import type { Post, CreatePostRequest, UpdatePostRequest } from '@/types/post'

export const postApi = {
  list(page = 1, pageSize = 10) {
    return client.get<ApiResponse<PaginatedData<Post>>>('/api/v1/posts', {
      params: { page, page_size: pageSize },
    })
  },

  get(id: number) {
    return client.get<ApiResponse<Post>>(`/api/v1/posts/${id}`)
  },

  create(data: CreatePostRequest) {
    return client.post<ApiResponse<Post>>('/api/v1/posts', data)
  },

  update(id: number, data: UpdatePostRequest) {
    return client.put<ApiResponse<Post>>(`/api/v1/posts/${id}`, data)
  },

  delete(id: number) {
    return client.delete<ApiResponse<null>>(`/api/v1/posts/${id}`)
  },
}
