pipeline {
    agent { label 'bm-inventory-subsystem' }
    environment {
        SERVICE = 'ocpmetal/assisted-service'
    }

    stages {
        stage('clear deployment') {
            steps {
                sh 'docker image prune -a -f'
                sh 'make clear-deployment'
            }
        }

        stage('Build') {
            steps {
                sh '''export PATH=$PATH:/usr/local/go/bin; make build-image'''
            }
        }

        stage('Deploy for subsystem') {
            steps {
                withCredentials([usernamePassword(credentialsId: 'dockerio_cred', passwordVariable: 'PASS', usernameVariable: 'USER')]) {
                    sh '''docker login docker.io -u $USER -p $PASS'''
                }
                sh '''export PATH=$PATH:/usr/local/go/bin; make jenkins-deploy-for-subsystem'''
                sh '''# Dump pod statuses;kubectl get pods -A'''
            }
        }

        stage('Subsystem-test') {
            steps {
                sh '''export PATH=$PATH:/usr/local/go/bin;make subsystem-run'''
            }
        }

        stage('clear deployment after subsystem test') {
            steps {
                sh 'make clear-deployment'
            }
        }

        stage('publish images on push to master') {
            when {
                branch 'master'
            }

            steps {
                withCredentials([usernamePassword(credentialsId: 'ocpmetal_cred', passwordVariable: 'PASS', usernameVariable: 'USER')]) {
                    sh '''docker login quay.io -u $USER -p $PASS'''
                }

                sh '''docker tag  ${SERVICE} quay.io/ocpmetal/assisted-service:latest'''
                sh '''docker tag  ${SERVICE} quay.io/ocpmetal/assisted-service:${GIT_COMMIT}'''
                sh '''docker push quay.io/ocpmetal/assisted-service:latest'''
                sh '''docker push quay.io/ocpmetal/assisted-service:${GIT_COMMIT}'''
            }
        }
    }

    post {
        failure {
            echo 'Get assisted-service log'
            sh '''
            kubectl get pods -o=custom-columns=NAME:.metadata.name -A | grep assisted-service | xargs -I {} sh -c "kubectl logs {} -n  assisted-installer > test_dd.log"
            mv test_dd.log $WORKSPACE/assisted-service.log || true
            '''

            echo 'Get postgres log'
            sh '''kubectl  get pods -o=custom-columns=NAME:.metadata.name -A | grep postgres | xargs -I {} sh -c "kubectl logs {} -n  assisted-installer > test_dd.log"
            mv test_dd.log $WORKSPACE/postgres.log || true
            '''

            echo 'Get scality log'
            sh '''kubectl  get pods -o=custom-columns=NAME:.metadata.name -A | grep scality | xargs -I {} sh -c "kubectl logs {} -n  assisted-installer > test_dd.log"
            mv test_dd.log $WORKSPACE/scality.log || true
            '''

            echo 'Get createimage log'
            sh '''kubectl  get pods -o=custom-columns=NAME:.metadata.name -A | grep createimage | xargs -I {} sh -c "kubectl logs {} -n  assisted-installer > test_dd.log"
            mv test_dd.log $WORKSPACE/createimage.log || true
            '''
        }
    }
}
