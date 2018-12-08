require 'active_model'

class User
  include ActiveModel::Validations

  attr_accessor :name
  validates :name, presence: true
end
