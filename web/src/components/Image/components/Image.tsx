import NextImage, { ImageProps as NativeImageProps } from 'next/image'
import { useCallback, useState } from 'react'
import { ImageSkeleton } from './ImageSkeleton'

interface LoadingImageProps extends Omit<NativeImageProps, 'src' | 'alt'> {
  src?: string
  alt?: string
  isLoading: true
}

interface NonLoadingImageProps extends Omit<NativeImageProps, 'src' | 'alt'> {
  src: string
  alt: string
  isLoading?: false
}

type ImageProps = LoadingImageProps | NonLoadingImageProps

export function Image({
  src,
  alt,
  width,
  height,
  isLoading,
  ...props
}: ImageProps) {
  const [isImageLoading, setIsImageLoading] = useState(true)

  const handleLoadingComplete = useCallback(() => {
    setIsImageLoading(false)
  }, [])

  return (
    <div style={{ position: 'relative', maxWidth: width }}>
      {isLoading && isImageLoading && (
        <ImageSkeleton
          style={{
            position: 'absolute',
            zIndex: 'auto',
            top: '-2px',
          }}
          width={width}
          height={height}
        />
      )}
      {!isLoading && (
        <NextImage
          {...props}
          src={src}
          alt={alt}
          width={width}
          height={height}
          onLoadingComplete={handleLoadingComplete}
        />
      )}
    </div>
  )
}
