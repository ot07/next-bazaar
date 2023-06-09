import { Select as MantineSelect } from '@mantine/core'
import { useController } from 'react-hook-form'
import { NumberSelectProps } from '../types'
import { ErrorMessage } from './ErrorMessage'

export function NumberSelect({
  label,
  options,
  name,
  ...rest
}: NumberSelectProps) {
  const {
    field,
    fieldState: { error: fieldError },
    formState: { defaultValues },
  } = useController({ name })

  const error = fieldError ? (
    <ErrorMessage>{fieldError.message?.toString()}</ErrorMessage>
  ) : undefined

  const { value, onChange, ...restField } = field

  return (
    <MantineSelect
      id={name}
      styles={{ rightSection: { pointerEvents: 'none' } }}
      label={label}
      value={value === undefined ? '' : value.toString()}
      onChange={(value) =>
        onChange(
          value === '' ? undefined : Number(value) ?? defaultValues?.[name]
        )
      }
      allowDeselect
      error={error}
      dropdownComponent="div"
      {...rest}
      data={options.map((option) => ({
        label: option.toString(),
        value: option.toString(),
      }))}
      {...restField}
    />
  )
}
