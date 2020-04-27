rm -f /var/lib/bashible/bootstrap-token
rm -f /var/lib/bashible/ca.crt
rm -f /var/lib/bashible/cloud-provider-bootstrap-networks-*.sh
rm -f /var/lib/bashible/detect_bundle.sh

# safety for re-bootstrap, look into 050_reset_control_plane_on_configuration_change.sh.tpl
find /.kubeadm.checksum -mmin +120 -delete >/dev/null 2>&1 || true
