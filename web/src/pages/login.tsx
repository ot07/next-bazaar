import {
  Anchor,
  Paper,
  Title,
  Text,
  Container,
  Button,
  Center,
  Stack,
} from "@mantine/core";
import { MantineLogo } from "@mantine/ds";
import Link from "next/link";
import { z } from "zod";
import { useCallback, useState } from "react";
import { useAuth } from "@/features/auth";
import { PasswordInput, TextInput, useForm } from "@/components/Form";
import { zodResolver } from "@hookform/resolvers/zod";

const schema = z.object({
  email: z
    .string()
    .min(1, { message: "Required" })
    .email({ message: "Wrong Format" }),
  password: z.string().min(8, { message: "Minimum 8 characters" }),
});

export default function LoginPage() {
  const [isLoginButtonClicked, setIsLoginButtonClicked] = useState(false);

  const handleLoginError = useCallback(() => {
    setIsLoginButtonClicked(false);
  }, []);

  const { login } = useAuth({
    onLoginError: handleLoginError,
  });

  const [Form, methods] = useForm<{
    email: string;
    password: string;
  }>({
    resolver: zodResolver(schema),
    defaultValues: {
      email: "",
      password: "",
    },
    onSubmit: (data) => {
      setIsLoginButtonClicked(true);
      login(data);
    },
  });

  return (
    <Container size={420} my={40}>
      <Stack>
        <Center>
          <Link href={"/"}>
            <MantineLogo size={30} />
          </Link>
        </Center>

        <Title
          align="center"
          sx={{
            fontWeight: 900,
          }}
        >
          Welcome back!
        </Title>

        <Paper withBorder shadow="md" p={30} radius="md">
          <Form>
            <TextInput
              type="email"
              name="email"
              label="Email"
              placeholder="you@email.com"
            />
            <PasswordInput
              name="password"
              label="Password"
              placeholder="Your password"
              mt="md"
            />
            <Button
              type="submit"
              fullWidth
              mt="xl"
              disabled={isLoginButtonClicked}
            >
              Log in
            </Button>
          </Form>
        </Paper>

        <Text color="dimmed" size="sm" align="center" mt={5}>
          Do not have an account yet?{" "}
          <Link href="/register">
            <Anchor size="sm" component="button">
              Create account
            </Anchor>
          </Link>
        </Text>
      </Stack>
    </Container>
  );
}