import React from 'react'
import { Select, MenuItem, Paper, makeStyles } from '@material-ui/core'
import Button from '@material-ui/core/Button'
import { useForm, Controller } from 'react-hook-form'
import Container from '@material-ui/core/Container'
import Typography from '@material-ui/core/Typography'
import TextField from '@material-ui/core/TextField'
import Switch from '@material-ui/core/Switch'

const defaultValues = {
  provider: '',
}

const useStyles = makeStyles(theme => ({
  root: {
    display: 'flex',
    flexDirection: 'column',
    justifyContent: 'center',
    alignItems: 'center',
    padding: theme.spacing(2),

    '& .MuiTextField-root': {
      width: '300px',
    },
    '& .MuiSelect-root': {
      width: '276px',
    },
    '& .MuiButton-root': {
      margin: theme.spacing(2, 0),
    },
    '& .marginTop': {
      margin: theme.spacing(3, 0, 0, 0),
    },
  },
}))

function Home() {
  const classes = useStyles()

  const { handleSubmit, control, register } = useForm({ defaultValues })
  const onSubmit = data => {
    if (data.provider === '') {
      alert('required provider missing')
      return
    }

    const requestOptions = {
      method: 'POST',
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    }

    fetch('http://localhost:8080/shipa-gen', requestOptions).then(response => {
      if (response.status !== 200) {
        alert('Bad request!')
        return
      }

      const filename = response.headers
        .get('content-disposition')
        .split('filename="')[1]
        .split('"')[0]

      response.blob().then(blob => {
        const url = window.URL.createObjectURL(new Blob([blob]))
        const link = document.createElement('a')
        link.href = url
        link.setAttribute('download', filename)
        document.body.appendChild(link)
        link.click()
        link.parentNode.removeChild(link)
      })
    })
  }

  return (
    <div>
      <Container component="main" maxWidth="xs">
        <Typography
          color="textPrimary"
          gutterBottom
          variant="h4"
          align="center"
        >
          App config
        </Typography>

        <form className={classes.root}>
          <Paper style={{ padding: 16 }}>
            <Typography color="textPrimary" variant="h6" align="left">
              Description
            </Typography>
            <Controller
              render={({ field }) => (
                <Select {...field}>
                  <MenuItem value="terraform">Terraform</MenuItem>
                  <MenuItem value="crossplane">Crossplane</MenuItem>
                  <MenuItem value="pulumi">Pulumi</MenuItem>
                  <MenuItem value="ansible">Ansible</MenuItem>
                  <MenuItem value="github">GitHub Actions</MenuItem>
                  <MenuItem value="cloudformation">CloudFormation</MenuItem>
                  <MenuItem value="gitlab">GitLab</MenuItem>
                </Select>
              )}
              control={control}
              name="provider"
            />
            <TextField label="Application name" {...register('appName')} />
            <TextField label="Framework" {...register('framework')} />
            <TextField label="Team" {...register('team')} />
            <TextField label="Plan" {...register('plan')} />
            <TextField label="Tags" {...register('tags')} />

            <Typography
              color="textPrimary"
              variant="h6"
              align="left"
              className={'marginTop'}
            >
              Deployment
            </Typography>

            <TextField label="Image" {...register('image')} />
            <TextField label="Registry user" {...register('registryUser')} />
            <TextField
              label="Registry secret"
              {...register('registrySecret')}
            />
            <TextField label="Port" {...register('port')} />

            <Typography
              color="textPrimary"
              variant="h6"
              align="left"
              className={'marginTop'}
            >
              CNAME
            </Typography>
            <TextField label="CNAME" {...register('cname')} />
            <Typography
              color="textSecondary"
              variant="subtitle1"
              align="left"
              display="inline"
            >
              Encrypt
            </Typography>
            <Switch label="Encrypt" {...register('encrypt')} />

            <Typography
              color="textPrimary"
              variant="h6"
              align="left"
              className={'marginTop'}
            >
              Environment Variables
            </Typography>
            <TextField label="Name" {...register('envName')} />
            <TextField label="Value" {...register('envValue')} />
            <Typography
              color="textSecondary"
              variant="subtitle1"
              align="left"
              display="inline"
            >
              Norestart
            </Typography>
            <Switch label="Restart" {...register('norestart')} />
            <Typography
              color="textSecondary"
              variant="subtitle1"
              align="left"
              display="inline"
            >
              Private
            </Typography>
            <Switch label="Private" {...register('private')} />

            <Button
              color="primary"
              size="large"
              fullWidth
              variant="contained"
              onClick={handleSubmit(onSubmit)}
            >
              Generate
            </Button>
          </Paper>
        </form>
      </Container>
    </div>
  )
}

export default Home
