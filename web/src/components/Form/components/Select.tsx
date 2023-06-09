import { Select as MantineSelect } from '@mantine/core'
import { useController } from 'react-hook-form'
import { SelectProps } from '../types'
import { ErrorMessage } from './ErrorMessage'

export function Select({ label, options, name, ...rest }: SelectProps) {
  const {
    field,
    fieldState: { error: fieldError },
    formState: { defaultValues },
  } = useController({ name })

  const error = fieldError ? (
    <ErrorMessage>{fieldError.message?.toString()}</ErrorMessage>
  ) : undefined

  const { onChange, ...restField } = field

  return (
    <MantineSelect
      id={name}
      styles={{ rightSection: { pointerEvents: 'none' } }}
      label={label}
      onChange={(value) => onChange(value ?? defaultValues?.[name])}
      allowDeselect
      error={error}
      dropdownComponent="div"
      {...rest}
      data={options}
      {...restField}
    />
  )
}
