/**
 * Generated by orval v6.15.0 🍺
 * Do not edit manually.
 * Next Bazaar API
 * OpenAPI spec version: 0.0.1
 */
import type {
  MutationFunction,
  QueryFunction,
  QueryKey,
  UseMutationOptions,
  UseQueryOptions,
  UseQueryResult,
} from '@tanstack/react-query'
import { useMutation, useQuery } from '@tanstack/react-query'
import type { ErrorType } from '../../custom-axios-instance'
import { customAxiosInstance } from '../../custom-axios-instance'
import type {
  ApiErrorResponse,
  ApiMessageResponse,
  CartDomainAddProductRequest,
  CartDomainCartProductResponse,
  CartDomainCartProductsCountResponse,
  CartDomainDeleteProductRequest,
  CartDomainUpdateProductQuantityRequest,
} from '../../model'

// eslint-disable-next-line
type SecondParameter<T extends (...args: any) => any> = T extends (
  config: any,
  args: infer P
) => any
  ? P
  : never

/**
 * @summary Delete cart product
 */
export const deleteCartProducts = (
  cartDomainDeleteProductRequest: CartDomainDeleteProductRequest,
  options?: SecondParameter<typeof customAxiosInstance>
) => {
  return customAxiosInstance<void>(
    {
      url: `/cart-products`,
      method: 'delete',
      headers: { 'Content-Type': 'application/json' },
      data: cartDomainDeleteProductRequest,
    },
    options
  )
}

export const getDeleteCartProductsMutationOptions = <
  TError = ErrorType<ApiErrorResponse>,
  TContext = unknown
>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof deleteCartProducts>>,
    TError,
    { data: CartDomainDeleteProductRequest },
    TContext
  >
  request?: SecondParameter<typeof customAxiosInstance>
}): UseMutationOptions<
  Awaited<ReturnType<typeof deleteCartProducts>>,
  TError,
  { data: CartDomainDeleteProductRequest },
  TContext
> => {
  const { mutation: mutationOptions, request: requestOptions } = options ?? {}

  const mutationFn: MutationFunction<
    Awaited<ReturnType<typeof deleteCartProducts>>,
    { data: CartDomainDeleteProductRequest }
  > = (props) => {
    const { data } = props ?? {}

    return deleteCartProducts(data, requestOptions)
  }

  return { mutationFn, ...mutationOptions }
}

export type DeleteCartProductsMutationResult = NonNullable<
  Awaited<ReturnType<typeof deleteCartProducts>>
>
export type DeleteCartProductsMutationBody = CartDomainDeleteProductRequest
export type DeleteCartProductsMutationError = ErrorType<ApiErrorResponse>

export const useDeleteCartProducts = <
  TError = ErrorType<ApiErrorResponse>,
  TContext = unknown
>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof deleteCartProducts>>,
    TError,
    { data: CartDomainDeleteProductRequest },
    TContext
  >
  request?: SecondParameter<typeof customAxiosInstance>
}) => {
  const mutationOptions = getDeleteCartProductsMutationOptions(options)

  return useMutation(mutationOptions)
}
/**
 * @summary Get cart
 */
export const getCartProducts = (
  options?: SecondParameter<typeof customAxiosInstance>,
  signal?: AbortSignal
) => {
  return customAxiosInstance<CartDomainCartProductResponse[]>(
    { url: `/cart-products`, method: 'get', signal },
    options
  )
}

export const getGetCartProductsQueryKey = () => [`/cart-products`] as const

export const getGetCartProductsQueryOptions = <
  TData = Awaited<ReturnType<typeof getCartProducts>>,
  TError = ErrorType<ApiErrorResponse>
>(options?: {
  query?: UseQueryOptions<
    Awaited<ReturnType<typeof getCartProducts>>,
    TError,
    TData
  >
  request?: SecondParameter<typeof customAxiosInstance>
}): UseQueryOptions<
  Awaited<ReturnType<typeof getCartProducts>>,
  TError,
  TData
> & { queryKey: QueryKey } => {
  const { query: queryOptions, request: requestOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getGetCartProductsQueryKey()

  const queryFn: QueryFunction<Awaited<ReturnType<typeof getCartProducts>>> = ({
    signal,
  }) => getCartProducts(requestOptions, signal)

  return { queryKey, queryFn, ...queryOptions }
}

export type GetCartProductsQueryResult = NonNullable<
  Awaited<ReturnType<typeof getCartProducts>>
>
export type GetCartProductsQueryError = ErrorType<ApiErrorResponse>

export const useGetCartProducts = <
  TData = Awaited<ReturnType<typeof getCartProducts>>,
  TError = ErrorType<ApiErrorResponse>
>(options?: {
  query?: UseQueryOptions<
    Awaited<ReturnType<typeof getCartProducts>>,
    TError,
    TData
  >
  request?: SecondParameter<typeof customAxiosInstance>
}): UseQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const queryOptions = getGetCartProductsQueryOptions(options)

  const query = useQuery(queryOptions) as UseQueryResult<TData, TError> & {
    queryKey: QueryKey
  }

  query.queryKey = queryOptions.queryKey

  return query
}

/**
 * @summary Add product to cart
 */
export const postCartProducts = (
  cartDomainAddProductRequest: CartDomainAddProductRequest,
  options?: SecondParameter<typeof customAxiosInstance>
) => {
  return customAxiosInstance<ApiMessageResponse>(
    {
      url: `/cart-products`,
      method: 'post',
      headers: { 'Content-Type': 'application/json' },
      data: cartDomainAddProductRequest,
    },
    options
  )
}

export const getPostCartProductsMutationOptions = <
  TError = ErrorType<ApiErrorResponse>,
  TContext = unknown
>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof postCartProducts>>,
    TError,
    { data: CartDomainAddProductRequest },
    TContext
  >
  request?: SecondParameter<typeof customAxiosInstance>
}): UseMutationOptions<
  Awaited<ReturnType<typeof postCartProducts>>,
  TError,
  { data: CartDomainAddProductRequest },
  TContext
> => {
  const { mutation: mutationOptions, request: requestOptions } = options ?? {}

  const mutationFn: MutationFunction<
    Awaited<ReturnType<typeof postCartProducts>>,
    { data: CartDomainAddProductRequest }
  > = (props) => {
    const { data } = props ?? {}

    return postCartProducts(data, requestOptions)
  }

  return { mutationFn, ...mutationOptions }
}

export type PostCartProductsMutationResult = NonNullable<
  Awaited<ReturnType<typeof postCartProducts>>
>
export type PostCartProductsMutationBody = CartDomainAddProductRequest
export type PostCartProductsMutationError = ErrorType<ApiErrorResponse>

export const usePostCartProducts = <
  TError = ErrorType<ApiErrorResponse>,
  TContext = unknown
>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof postCartProducts>>,
    TError,
    { data: CartDomainAddProductRequest },
    TContext
  >
  request?: SecondParameter<typeof customAxiosInstance>
}) => {
  const mutationOptions = getPostCartProductsMutationOptions(options)

  return useMutation(mutationOptions)
}
/**
 * @summary Update cart product quantity
 */
export const putCartProducts = (
  cartDomainUpdateProductQuantityRequest: CartDomainUpdateProductQuantityRequest,
  options?: SecondParameter<typeof customAxiosInstance>
) => {
  return customAxiosInstance<ApiMessageResponse>(
    {
      url: `/cart-products`,
      method: 'put',
      headers: { 'Content-Type': 'application/json' },
      data: cartDomainUpdateProductQuantityRequest,
    },
    options
  )
}

export const getPutCartProductsMutationOptions = <
  TError = ErrorType<ApiErrorResponse>,
  TContext = unknown
>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof putCartProducts>>,
    TError,
    { data: CartDomainUpdateProductQuantityRequest },
    TContext
  >
  request?: SecondParameter<typeof customAxiosInstance>
}): UseMutationOptions<
  Awaited<ReturnType<typeof putCartProducts>>,
  TError,
  { data: CartDomainUpdateProductQuantityRequest },
  TContext
> => {
  const { mutation: mutationOptions, request: requestOptions } = options ?? {}

  const mutationFn: MutationFunction<
    Awaited<ReturnType<typeof putCartProducts>>,
    { data: CartDomainUpdateProductQuantityRequest }
  > = (props) => {
    const { data } = props ?? {}

    return putCartProducts(data, requestOptions)
  }

  return { mutationFn, ...mutationOptions }
}

export type PutCartProductsMutationResult = NonNullable<
  Awaited<ReturnType<typeof putCartProducts>>
>
export type PutCartProductsMutationBody = CartDomainUpdateProductQuantityRequest
export type PutCartProductsMutationError = ErrorType<ApiErrorResponse>

export const usePutCartProducts = <
  TError = ErrorType<ApiErrorResponse>,
  TContext = unknown
>(options?: {
  mutation?: UseMutationOptions<
    Awaited<ReturnType<typeof putCartProducts>>,
    TError,
    { data: CartDomainUpdateProductQuantityRequest },
    TContext
  >
  request?: SecondParameter<typeof customAxiosInstance>
}) => {
  const mutationOptions = getPutCartProductsMutationOptions(options)

  return useMutation(mutationOptions)
}
/**
 * @summary Get cart products count
 */
export const getCartProductsCount = (
  options?: SecondParameter<typeof customAxiosInstance>,
  signal?: AbortSignal
) => {
  return customAxiosInstance<CartDomainCartProductsCountResponse>(
    { url: `/cart-products/count`, method: 'get', signal },
    options
  )
}

export const getGetCartProductsCountQueryKey = () =>
  [`/cart-products/count`] as const

export const getGetCartProductsCountQueryOptions = <
  TData = Awaited<ReturnType<typeof getCartProductsCount>>,
  TError = ErrorType<ApiErrorResponse>
>(options?: {
  query?: UseQueryOptions<
    Awaited<ReturnType<typeof getCartProductsCount>>,
    TError,
    TData
  >
  request?: SecondParameter<typeof customAxiosInstance>
}): UseQueryOptions<
  Awaited<ReturnType<typeof getCartProductsCount>>,
  TError,
  TData
> & { queryKey: QueryKey } => {
  const { query: queryOptions, request: requestOptions } = options ?? {}

  const queryKey = queryOptions?.queryKey ?? getGetCartProductsCountQueryKey()

  const queryFn: QueryFunction<
    Awaited<ReturnType<typeof getCartProductsCount>>
  > = ({ signal }) => getCartProductsCount(requestOptions, signal)

  return { queryKey, queryFn, ...queryOptions }
}

export type GetCartProductsCountQueryResult = NonNullable<
  Awaited<ReturnType<typeof getCartProductsCount>>
>
export type GetCartProductsCountQueryError = ErrorType<ApiErrorResponse>

export const useGetCartProductsCount = <
  TData = Awaited<ReturnType<typeof getCartProductsCount>>,
  TError = ErrorType<ApiErrorResponse>
>(options?: {
  query?: UseQueryOptions<
    Awaited<ReturnType<typeof getCartProductsCount>>,
    TError,
    TData
  >
  request?: SecondParameter<typeof customAxiosInstance>
}): UseQueryResult<TData, TError> & { queryKey: QueryKey } => {
  const queryOptions = getGetCartProductsCountQueryOptions(options)

  const query = useQuery(queryOptions) as UseQueryResult<TData, TError> & {
    queryKey: QueryKey
  }

  query.queryKey = queryOptions.queryKey

  return query
}
